/* @refresh reload */
import { render } from 'solid-js/web';

import 'flowbite';
import './index.css';
import App from './app';

const root = document.getElementById('root');

render(() => <App />, root!);
