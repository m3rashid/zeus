import { lazy } from 'solid-js';
import { Router, Route } from '@solidjs/router';
import { ColorModeProvider } from '@kobalte/core';

const Home = lazy(() => import('./pages/home'));
const Auth = lazy(() => import('./pages/auth'));

const App = () => {
  return (
    <ColorModeProvider>
      <Router>
        <Route path='/' component={Home} />
        <Route path='/auth' component={Auth} />
      </Router>
    </ColorModeProvider>
  );
};

export default App;
