import React from 'react';
import ComponentCreator from '@docusaurus/ComponentCreator';

export default [
  {
    path: '/usm/blog',
    component: ComponentCreator('/usm/blog', '511'),
    exact: true
  },
  {
    path: '/usm/blog/2025/08/20/welcome',
    component: ComponentCreator('/usm/blog/2025/08/20/welcome', 'd48'),
    exact: true
  },
  {
    path: '/usm/blog/archive',
    component: ComponentCreator('/usm/blog/archive', '7ce'),
    exact: true
  },
  {
    path: '/usm/docs',
    component: ComponentCreator('/usm/docs', 'aae'),
    routes: [
      {
        path: '/usm/docs',
        component: ComponentCreator('/usm/docs', '0df'),
        routes: [
          {
            path: '/usm/docs',
            component: ComponentCreator('/usm/docs', '182'),
            routes: [
              {
                path: '/usm/docs/',
                component: ComponentCreator('/usm/docs/', 'b6f'),
                exact: true
              },
              {
                path: '/usm/docs/ci-cd-recipes',
                component: ComponentCreator('/usm/docs/ci-cd-recipes', 'de5'),
                exact: true,
                sidebar: "docs"
              },
              {
                path: '/usm/docs/developer-guide',
                component: ComponentCreator('/usm/docs/developer-guide', '74f'),
                exact: true,
                sidebar: "docs"
              },
              {
                path: '/usm/docs/enterprise-features',
                component: ComponentCreator('/usm/docs/enterprise-features', '188'),
                exact: true,
                sidebar: "docs"
              },
              {
                path: '/usm/docs/format-spec',
                component: ComponentCreator('/usm/docs/format-spec', 'f3e'),
                exact: true,
                sidebar: "docs"
              },
              {
                path: '/usm/docs/migration-from-dotenv',
                component: ComponentCreator('/usm/docs/migration-from-dotenv', '17c'),
                exact: true,
                sidebar: "docs"
              },
              {
                path: '/usm/docs/quickstart',
                component: ComponentCreator('/usm/docs/quickstart', '6c5'),
                exact: true,
                sidebar: "docs"
              },
              {
                path: '/usm/docs/sdk-guides',
                component: ComponentCreator('/usm/docs/sdk-guides', 'bea'),
                exact: true,
                sidebar: "docs"
              },
              {
                path: '/usm/docs/security-model',
                component: ComponentCreator('/usm/docs/security-model', 'a72'),
                exact: true,
                sidebar: "docs"
              }
            ]
          }
        ]
      }
    ]
  },
  {
    path: '*',
    component: ComponentCreator('*'),
  },
];
