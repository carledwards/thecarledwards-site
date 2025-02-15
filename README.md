# Custom Hugo Theme

This site uses a customized version of the M10C theme. The original theme is kept as a reference while custom modifications are made in the root directories.

## Directory Structure for Customization

```
.
├── assets/
│   └── css/                    # Custom CSS/SCSS files
│       ├── main.scss          # Main stylesheet (override theme styles here)
│       ├── _base.scss         # Base styles
│       └── components/        # Component-specific styles
├── layouts/
│   ├── _default/              # Default layout templates
│   │   ├── baseof.html       # Base template
│   │   └── single.html       # Single page template
│   └── partials/             # Partial templates
```

## How to Customize

1. **Styling Changes**
   - Modify files in `assets/css/` to change the look and feel
   - Main color scheme can be adjusted in `config.toml` under `[params.style]`
   - Component-specific styles are in `assets/css/components/`

2. **Layout Changes**
   - Edit templates in `layouts/_default/` to modify page structure
   - `baseof.html` is the main template that wraps all pages
   - `single.html` controls the layout of individual posts/pages

3. **Adding New Features**
   - Create new partial templates in `layouts/partials/`
   - Add new SCSS files in `assets/css/components/`
   - Reference new partials in existing templates

## Original Theme Reference
The original M10C theme is kept in `themes/hugo-theme-m10c/` for reference. Any file in the root `layouts/` or `assets/` directory will override the corresponding theme file.
