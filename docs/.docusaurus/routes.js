import React from 'react';
import ComponentCreator from '@docusaurus/ComponentCreator';

export default [
  {
    path: '/usm/blog',
    component: ComponentCreator('/usm/blog', '592'),
    exact: true
  },
  {
    path: '/usm/blog/2025/08/20/welcome',
    component: ComponentCreator('/usm/blog/2025/08/20/welcome', '0a4'),
    exact: true
  },
  {
    path: '/usm/blog/archive',
    component: ComponentCreator('/usm/blog/archive', 'cbb'),
    exact: true
  },
  {
    path: '/usm/docs',
    component: ComponentCreator('/usm/docs', '662'),
    routes: [
      {
        path: '/usm/docs/',
        component: ComponentCreator('/usm/docs/', 'b6e'),
        exact: true
      },
      {
        path: '/usm/docs/ci-cd-recipes',
        component: ComponentCreator('/usm/docs/ci-cd-recipes', '832'),
        exact: true,
        sidebar: "docs"
      },
      {
        path: '/usm/docs/developer-guide',
        component: ComponentCreator('/usm/docs/developer-guide', '9da'),
        exact: true,
        sidebar: "docs"
      },
      {
        path: '/usm/docs/enterprise-features',
        component: ComponentCreator('/usm/docs/enterprise-features', '768'),
        exact: true,
        sidebar: "docs"
      },
      {
        path: '/usm/docs/format-spec',
        component: ComponentCreator('/usm/docs/format-spec', '07e'),
        exact: true,
        sidebar: "docs"
      },
      {
        path: '/usm/docs/migration-from-dotenv',
        component: ComponentCreator('/usm/docs/migration-from-dotenv', 'b1c'),
        exact: true,
        sidebar: "docs"
      },
      {
        path: '/usm/docs/quickstart',
        component: ComponentCreator('/usm/docs/quickstart', '2d9'),
        exact: true,
        sidebar: "docs"
      },
      {
        path: '/usm/docs/sdk-guides',
        component: ComponentCreator('/usm/docs/sdk-guides', '44f'),
        exact: true,
        sidebar: "docs"
      },
      {
        path: '/usm/docs/security-model',
        component: ComponentCreator('/usm/docs/security-model', '6f2'),
        exact: true,
        sidebar: "docs"
      }
    ]
  },
  {
    path: '*',
    component: ComponentCreator('*'),
  },
];
