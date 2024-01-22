import { Route, Router } from '@solidjs/router';
import { lazy } from 'solid-js';

const Auth = lazy(() => import('@web/pages/auth'));
const ChooseUser = lazy(() => import('./pages/chooseUser'));

const App = () => {
  return (
    <Router>
      <Route path='/auth/login' component={Auth} />
      <Route path='/auth/register' component={Auth} />
      <Route path='/auth/choose-user' component={ChooseUser} />
    </Router>
  );
};

export default App;
