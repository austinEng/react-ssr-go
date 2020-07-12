import * as React from 'react';
import { hydrate } from 'react-dom';
import App from './App';

const props = (() => {
  const stateHolder = (window as { __SSR_STATE__?: string });
  const ssrState = stateHolder.__SSR_STATE__;
  if (ssrState) {
    delete stateHolder.__SSR_STATE__;
    return JSON.parse(ssrState);
  }
  return {};
})();

hydrate(<App {...props} />, document.getElementById('root'));