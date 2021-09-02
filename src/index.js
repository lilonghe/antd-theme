import 'react-hot-loader';
import React from 'react';
import ReactDOM from 'react-dom';
import App from './app';

if (process.env.NODE_ENV !== 'production') {
    import('./themes/red.less')
}

ReactDOM.render(<App />, document.querySelector('#root'));

