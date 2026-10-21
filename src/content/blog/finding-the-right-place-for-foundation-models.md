---
title: "Finding the Right Place for Apple's OS 27 Foundation Models"
kind: "Project update"
date: 2026-09-05
description: "OS 27 brings substantially more capable models to the device, with Private Cloud Compute available when a feature needs more. I built a workbench to understand where each model fits—and what that changes about application architecture."
tags: ["Apple Foundation Models", "On-device AI", "Private Cloud Compute", "Architecture", "Evaluation", "iOS"]
---

Apple's OS 27 releases make a substantially more capable generation of
Foundation Models available to applications.

Some requests can run entirely on the device, including tasks that interpret
text and images. More demanding work can move to a larger model through Private
Cloud Compute. Both are available through the same Foundation Models framework.

That creates a more interesting question than how to call the API:

> Where should the intelligence for a feature actually run?

Running locally can reduce latency, avoid recurring inference costs, preserve
privacy, work without a network connection, and remove the need to build a
backend for certain kinds of applications. Private Cloud Compute can provide
additional context and reasoning when the feature genuinely needs it.

The right answer will be different for every product—and it cannot be determined
from a single successful prompt.

I built the **Apple Foundation Models Workbench** to understand how these models
behave against specific product requirements: what they can interpret, how well
they follow instructions, how consistent their responses are, how long they take,
and when moving from an on-device model to Private Cloud Compute makes a
meaningful difference.

One small image experiment quickly showed why that evaluation matters. I gave the
models a photograph of two dogs and a cat and asked which animal was looking at
the camera.

![The photograph and the run configuration on iOS — two dogs and a cat on a lawn](/images/foundation-models-workbench/image-run-setup.png)

*The photograph the models were given, with the instruction to only report on
animals. Which one is looking at the camera?*

![The same question run three times on-device and three times through Private Cloud Compute](/images/foundation-models-workbench/consistency-vs-correctness.png)

The on-device model returned the same answer three times. It was also wrong all
three times.

Private Cloud Compute recognized more of the scene, but it still failed to
isolate the animal looking at the camera.

Consistency was not correctness—but that was only one part of what the experiment
revealed. More broadly, it showed that choosing where a feature runs requires
understanding the behavior and constraints of each model, not simply selecting
the most capable one.

## Why placement is a real decision

Apple's third generation of Foundation Models puts AFM 3 Core and AFM 3 Core
Advanced on the device, alongside larger models reachable through Private Cloud
Compute. AFM 3 Core Advanced is the most capable of the on-device options, is
natively multimodal, and uses a sparse architecture so it can be more capable
without lighting up every parameter on every request.
[Apple describes the family and the architecture here](https://machinelearning.apple.com/research/introducing-third-generation-of-apple-foundation-models).

Private Cloud Compute offers a larger model through the same programming model—but
it needs the network and is subject to availability, entitlement and usage limits.
Apple's own guidance is to evaluate with the on-device model first and reach for
PCC when the feature genuinely needs it.
[Those tradeoffs are documented here](https://developer.apple.com/documentation/foundationmodels/adding-server-side-intelligence-with-private-cloud-compute).

Those constraints are what make this an architectural decision rather than a
configuration value. A feature that escalates to PCC has to behave sensibly when
the network is gone, when the entitlement is missing, or when usage limits are
reached. Deciding that after the feature is built is expensive.

## Where this started

The idea came from my outdoor security cameras.

The cameras already generate motion and object-detection events, but those
triggers carry very little context. Knowing that a camera detected an animal is
mildly useful. Knowing *which* animal, what it appears to be doing, and whether
the event is worth a notification would be considerably more useful.

So the question was whether an Apple Foundation Model could take a camera frame
and produce something structured enough for another system to act on—sketching
roughly toward this, rather than a generic animal-detected event:

```json
{
  "animals": ["cat"],
  "activity": "walking toward the side gate",
  "attention": "not looking at the camera",
  "notificationPriority": "low"
}
```

Keeping that on the device would cut latency, avoid shipping images to someone
else's service, and remove a recurring inference bill. All of which matters only
if the output is good enough to build on.

## What the workbench does

macOS 27 ships an `fm` command-line tool that already makes it possible to poke
at these models without rebuilding an Xcode project every time—interactive
sessions, single prompts, image input, model selection, structured output.
[Apple introduced it at WWDC](https://developer.apple.com/videos/play/wwdc2026/334/).
The workbench takes the same idea and points it at repeatable product
experiments, across three surfaces: a local HTTP server you can hit with `curl`,
a native macOS app, and a native iOS app.

The part that matters is running one prompt several times and comparing what
comes back. **The agreement score is just the share of runs that produced the
identical response**—three matching answers out of three is `1.00 3/3`; a two-one
split is `0.67 2/3` against `0.33 1/3`.

![The macOS image tab showing a prompt split two ways across three samples](/images/foundation-models-workbench/agreement-split.png)

*A different experiment from the one above—no system instruction here, and the
schema asks for "every distinct thing visible in the image", which is why one run
volunteers `grass`. Two runs out of three agreed.*

Three samples is a small number, and worth being honest about: it can tell you a
model is unstable, but it cannot tell you a model is stable. It is a smoke test,
not a measurement.

## Consistency is not correctness

The animal experiment is worth looking at in detail, because both models failed
in different and instructive ways.

The system instruction was *"only report on animals"*, the schema asked for a list
of subjects, and the question was *"What animal is looking at the camera?"* Three
runs on the device, three through PCC.

On the device, all three runs came back identical:

```json
{
  "subjects": ["cat"]
}
```

A perfect `1.00 3/3` agreement score. It is also wrong. The cat is walking with
its head down, looking at the grass. The white-and-brown dog in the middle of the
frame is the one staring straight into the lens.

Private Cloud Compute split two ways: two runs returned the cat and the dog, one
returned the same two animals in the opposite order. It recognized more of the
scene, but neither response actually answered the question correctly. The
workbench reported `3× · 2,861 ms` for the on-device run and `3× · 3,583 ms` for
PCC.

Two things fall out of that.

**A high agreement score says nothing about correctness.** It says the model
behaved consistently across those runs. Proving the answer right still needs known
examples and expected results.

**Comparing whole JSON objects surfaces differences the product may not care
about.** `["cat", "dog"]` and `["dog", "cat"]` counted as two different answers,
because the workbench normalizes key order but not the order of items inside a
list. For this feature that distinction is noise. If order does not matter to the
product, the evaluator should sort those arrays before comparing—and that is a
decision about the product, not about the model.

## Instructions are a separate contract

Consistency is one way a model disappoints you. Ignoring your instructions is
another, and it shows up in plain text prompts too.

The system prompt for these runs was explicit:

> You are a kitchen helper. The user tells you what they want to cook and for
> whom; you work out what they need. Be specific about ingredients: give
> quantities, the form (fresh, canned, frozen, dried), and details that matter
> such as unsweetened, low-fat, or gluten-free.

Same instructions, same JSON schema, same request—a Thanksgiving dessert for six.
Both models returned valid structured output and both picked pumpkin pie.

![AFM 3 Core returning four ingredients with no forms](/images/foundation-models-workbench/kitchen-afm3-core.png)

*AFM 3 Core: four ingredients, no forms at all.*

![Private Cloud Compute returning nine ingredients with forms and dietary detail](/images/foundation-models-workbench/kitchen-pcc.png)

*Private Cloud Compute: nine ingredients, with `canned`, `fresh`, `low-fat`,
`dried`, `unsweetened` and `gluten-free if needed`.*

The instructions asked for the form of each ingredient; PCC supplied it and AFM 3
Core did not. Schema compliance and instruction compliance are separate things,
and only one of them is enforced for you by guided generation.

That is still not a verdict on which model is better. A shopping assistant needs
the quantities and forms. A suggestion surface might prefer speed and variety. A
feature that has to work on a plane needs the one that runs with the network off.
Which model is more capable is rarely the useful question.

Worth noting for anyone planning around context limits: the on-device session
reported `319 / 8,192 tokens`, while the PCC session showed *"Token counting is
not available for Private Cloud Compute sessions."* Different models, different
operational visibility.

## The backend does not have to be where the intelligence runs

On-device inference does not necessarily remove the backend. It changes the
backend's role.

Instead of sending every image or piece of text to a server for inference, an
application could perform the interpretation locally. A backend could then
coordinate that behavior by providing:

- versioned system instructions
- feature-specific prompt configurations
- output schemas
- model-selection rules
- feature flags and staged rollouts
- evaluation examples and expected results
- thresholds for escalating a request to a remote model

This would allow a team to improve how an intelligent feature behaves without
embedding every decision permanently in the application or waiting for a new App
Store release.

The application could also return carefully designed signals about performance:
whether the model produced valid output, whether the result was accepted or
corrected, which prompt version was used, how long it took, and whether the
request required escalation. That feedback could help answer whether a prompt
change improved the feature across real conditions—not just the examples used
during development.

Privacy matters here. An application should not automatically send private model
inputs or complete responses back to a server. In many cases, aggregate outcomes
and explicit user feedback may provide enough signal without transmitting the
content that was interpreted locally.

The backend no longer has to be the place where intelligence happens. It can be
the system that helps configure, measure, and improve intelligence running
elsewhere.

## What local intelligence could unlock

The other opportunity is applications that do not need a backend at all.

A sufficiently capable local model can give an application flexible behavior
without requiring developers to anticipate and script every possible input. It
can interpret natural language, extract structured information, classify content,
reason about an image, or choose between a constrained set of actions.

That could support applications which:

- interpret a photograph and organize it locally
- turn natural-language instructions into structured actions
- classify personal content without uploading it
- add context to sensor or camera events
- create lightweight automations from a user's description
- adapt an interface based on what the user is trying to accomplish
- continue providing useful behavior without a network connection

These are not unrestricted autonomous agents. They are focused product features in
which a local model provides interpretation and flexibility that would otherwise
require a collection of rules—or a recurring call to a hosted model.

That changes the threshold for adding intelligent behavior to an application. A
developer can create something useful, private, and responsive without first
operating an inference service.

## Finding the boundary

The workbench is my way of finding the boundary between those approaches.

The next step for the camera experiment is a representative set of images with
expected results. That will let me measure more than agreement:

- whether the model identifies the relevant subject
- how reliably it follows the instructions
- which scenes produce incorrect or incomplete answers
- whether a prompt revision improves the whole set
- how latency changes with the model and image size
- when Private Cloud Compute provides a meaningful improvement
- what the feature should do when the preferred model is unavailable

The result might be entirely local. It might use the device for an initial
interpretation and escalate uncertain cases. It might need a remote model for the
complete task. The point is to understand that before building the rest of the
product around an assumption.

OS 27 opens up a broader category of applications: some with no backend, some with
a backend that coordinates local intelligence, and some that move between local
and remote models depending on the request.

The interesting work now is learning what belongs where.

The workbench is open source—the [project page](/projects/#apple-intelligence-foundation-server)
has more detail, and the code is
[on GitHub](https://github.com/carledwards/apple-intelligence-foundation-server).

I plan to write separately about the security-camera implementation, the
evaluation dataset, and how all of this shapes the architecture around the model.
