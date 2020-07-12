import * as React from 'react';
import * as PropTypes from 'prop-types';
import { Fragment } from 'react';
import { Helmet } from 'react-helmet';

import ItemList from './components/ItemList';

type AppProps = {
  pageTitle: string,
  items: string[],
}

function App({ pageTitle, items }: AppProps) {
  return (
    <Fragment>
      <Helmet>
        <title>{pageTitle}</title>
      </Helmet>
      <h1>Hello {pageTitle}!</h1>
      <ItemList items={items} />
    </Fragment>
  )
}

App.propTypes = {
  pageTitle: PropTypes.string,
  items: PropTypes.arrayOf(PropTypes.string),
};

App.defaultProps = {
  pageTitle: 'App',
  items: [],
};

export default App;