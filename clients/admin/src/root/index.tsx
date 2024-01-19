import { ParentComponent } from 'solid-js';
import Header from './header';

const RootContainer: ParentComponent = (props) => {
  return (
    <>
      <Header />
      {props.children}
    </>
  );
};

export default RootContainer;
