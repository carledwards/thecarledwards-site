---
title: "WiFi on DIY"
date: 2025-03-08
description: "Explore the challenges and creative solutions for integrating WiFi into DIY projects, from using the ESP32 to navigating protocols like Matter and Apple's MFI."
---

For my personal projects, I really enjoy building solutions that run on hardware like the ESP32. It has WiFi and Bluetooth capabilities and works well with MicroPython. Some of these projects benefit from an internet connection—for example, using NTP (Network Time Protocol) to obtain the current date and time or hosting a self-contained web server. When choosing hardware, I tend to prioritize options with built-in WiFi. It gives me the flexibility to expand a project’s capabilities, even when I’m still figuring out its final goal.

But getting a DIY project connected to the internet easily can be cumbersome.

Our mobile devices already store WiFi credentials (SSID and password). With some commercial devices, your phone can transfer these credentials with a single tap. I saw this in action while setting up my Elgato Ring Light. The device appeared as a special entry in my iPhone’s WiFi settings, and with one press, it joined my network.

As a DIYer, I’d love to use this feature myself. And technically, I can—but there are significant hurdles.

## Apple’s MFi Program: A Costly Gatekeeper

For an Apple/iOS solution, I’d have to become a registered vendor in Apple’s MFi (Made for iPhone) program. That would allow me to purchase Apple’s special chip, which grants my device access to this trusted ecosystem. But for DIY projects, this route isn’t practical due to high costs and stringent requirements.

## Matter Protocol: An Industry-Locked Option

Another approach is the Matter protocol, which is gaining traction for home automation. Developed by the Connectivity Standards Alliance, Matter aims to reduce vendor lock-in for smart home devices. However, participation requires being a registered company and embedding a unique, signed digital certificate in your device to send WiFi credentials. Again, this certification process is out of reach for most DIYers.

## Alternative DIY Methods

Since these solutions are unavailable to hobbyists, here are some workarounds:
	1.	QR Code Scanning
If the device has a camera, it can scan a QR code containing the WiFi SSID and password. This is usually done through a custom mobile app: you enter your credentials, generate a QR code, and display it for the device to scan. If it fails to connect, the device might blink an LED or emit a sound to indicate an error.
	2.	Bluetooth Credential Transfer
A more robust method is using Bluetooth. A custom mobile app can send the WiFi credentials over Bluetooth, allowing the device to connect without requiring manual input. This approach is more seamless and avoids potential security risks associated with QR code scanning.

## The Password Problem

For me, there’s a familiar pattern: I don’t remember my WiFi password, so I look it up in my router’s app, copy it to the clipboard, and switch back to my DIY app to paste it in. However, storing passwords in the clipboard is a security risk. While operating systems are improving clipboard security, it’s still not universally trusted. As a result, many people create weak WiFi passwords just so they’re easier to remember and type in.

## Locked WiFi Sharing: A Missed Opportunity for DIY

All I really want is for my ESP32 to quickly connect to WiFi. However, existing solutions—whether Apple’s proprietary ecosystem or Matter-enabled devices (like Apple Home, Google Home, Amazon Alexa, and Samsung SmartThings)—restrict who can access WiFi credentials. This means we must trust these companies to ensure that certified devices aren’t misusing our home network credentials or monitoring internet traffic.

The reality is, unless you actively monitor your network traffic, you have no way of knowing what these devices are capturing and sharing.

## A Call for Open WiFi Sharing Protocols

I wish iOS and Android would open up these WiFi-sharing protocols to DIYers. It would make it easier to create unique and useful gadgets, share them with loved ones, and avoid frustrating WiFi setup processes.

## A Better Way to Share WiFi Credentials

Here are a few ideas that could improve WiFi sharing, especially for setting up devices for others:
	•	Allow users to share WiFi credentials by generating a temporary pass with an expiration time.
	•	Restrict shared credentials so they cannot be forwarded to others.
	•	Never expose the master WiFi password to third-party devices (e.g., smart switches, thermostats, cameras). Instead, create a sandboxed network where their traffic can be monitored and controlled.

For now, my preferred solution is using a custom app with Bluetooth-based credential transfer. It’s the easiest and most secure option available to me.

Until something changes, I’ll keep rolling my own solutions—and hope my custom app gets approved in the app stores so I can provide a better experience for family and friends.
