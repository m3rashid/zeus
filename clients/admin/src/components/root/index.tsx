import { ParentComponent } from 'solid-js';
import Header from './header';
import SecondaryHeader from './headerSecondary';

const RootContainer: ParentComponent = (props) => {
  return (
    <>
      <Header />
      <SecondaryHeader />
      {props.children}
    </>
  );
};

export default RootContainer;
