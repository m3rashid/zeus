/* @refresh reload */
import { render } from 'solid-js/web';

import '@web/index.css';
import App from '@web/app';

const root = document.getElementById('root');

render(() => <App />, root!);
