/* @refresh reload */
import { render } from 'solid-js/web';
import '@zeus/components/index.css';

// import './index.css';
import App from './app';

const root = document.getElementById('root');

render(() => <App />, root!);
