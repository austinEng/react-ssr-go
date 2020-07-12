import * as React from 'react';

function ItemList(props: { items?: string[]}) {
  return (
    <div>
      <h2>Items</h2>
      <ul>
        {(props.items || []).map((item, i) => (
          <li key={i}>{item}</li>
        ))}
      </ul>
    </div>
  )
}

export default ItemList;
