# Docusaurus Documentation Site

This directory contains the configuration for the Docusaurus documentation site.

## Development

```bash
npm start
```

## Build

```bash
npm run build
```

## Deployment

The documentation site is automatically deployed to GitHub Pages whenever changes are pushed to the `main` branch. The site is available at:

https://universal-secrets-manager.github.io/usm/

## Adding New Documentation

To add new documentation pages:

1. Create a new Markdown file in the `docs` directory
2. Add the file to the `sidebars.js` configuration
3. Update the `docs/index.md` file to include a link to the new page