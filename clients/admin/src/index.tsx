/* @refresh reload */
import { render } from 'solid-js/web';

import '@admin/index.css';
import App from '@admin/app';

const root = document.getElementById('root');

render(() => <App />, root!);
