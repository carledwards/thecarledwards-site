+++
title = "Setting Up My Blog with Hugo and Cloudflare Pages"
date = "2024-02-15"
tags = ["hugo", "cloudflare", "setup"]
description = "A step-by-step guide on setting up a Hugo blog with Cloudflare Pages, including local development, deployment, and custom domain configuration."
+++

I recently decided to create a blog and wanted to share my setup process. I already had a domain name registered and chose to use Cloudflare Pages (their free hosting service) along with the Hugo framework for the implementation.

<!--more-->

> Note: While I went with Hugo, Cloudflare Pages supports many other frameworks including Next.js, Gatsby, Jekyll, Eleventy, and more. You can find the full list in their [documentation](https://developers.cloudflare.com/pages/framework-guides/).

Prerequisites:
- A registered domain name
- macOS (though the steps are similar for other operating systems)
- Basic comfort with command line operations
- GitHub account
- Cloudflare account (free tier is sufficient)

Here's the technical guide on how I set everything up:

1. First, create an empty repository on GitHub and clone it:
```bash
git clone https://github.com/yourusername/your-repo-name.git
cd your-repo-name
```

2. Install Hugo using Homebrew:
```bash
brew install hugo
```

3. Initialize the Hugo site (the `--force` flag is needed since the directory isn't empty due to git):
```bash
hugo new site . --force
```

4. At this point, you'll notice a `hugo.toml` file was created. Rename it to `config.toml`:
```bash
mv hugo.toml config.toml
```

5. Choose and set up a theme. I browsed the Hugo themes and found one I liked. Add it as a git submodule:
```bash
git submodule add https://github.com/theme-author/theme-name.git themes/theme-name
```

6. Test the site locally:
```bash
hugo server -D
```

7. Once everything looks good locally, commit and push to GitHub:
```bash
git add .
git commit -m "Initial Hugo site setup"
git push origin main
```

8. Set up Cloudflare Pages:
   - Go to Cloudflare dashboard
   - Select "Workers & Pages"
   - Create a new Pages project
   - Connect your GitHub repository
   - Important: Add an environment variable named `HUGO_VERSION` with your local Hugo version
     ```bash
     # Check your Hugo version with:
     hugo version
     ```

9. For custom domain setup:
   - In Cloudflare Pages, go to your project
   - Select "Custom Domains"
   - Follow the prompts to set up your domain
   - Update your domain registrar's nameservers to Cloudflare's
   - Wait for propagation (took about 20 minutes in my case)
   - On macOS, flush your DNS cache:
     ```bash
     sudo dscacheutil -flushcache; sudo killall -HUP mDNSResponder
     ```

That's it! Now every push to the main branch automatically triggers a new deployment through Cloudflare Pages. The setup might seem involved, but it creates a robust, fast, and easily maintainable blog infrastructure.
