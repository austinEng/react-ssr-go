import * as React from 'react';
import { renderToString } from 'react-dom/server';
import { Helmet } from 'react-helmet';
import App from '../src/App';

export const Index = (params: string | undefined) => {
  const props = params ? JSON.parse(params) : {};
  const app = renderToString(<App {...props} />);
  const helmetData = Helmet.renderStatic();

  return (
`<!doctype html>
<html>
  <head>
    ${helmetData.title.toString()}
    ${helmetData.meta.toString()}
    <script async src="/js/app.js"></script>
    <script>
      window.__SSR_STATE__ = ${JSON.stringify(params)};
    </script>
  </head>
  <body>
    <div id="root">${app}</div>
  </body>
</html>`
  );
};