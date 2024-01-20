import RootContainer from '@admin/components/root';
import { Input } from '@admin/components/ui/input';
import { Label } from '@admin/components/ui/label';

const HomePage = () => {
  return (
    <RootContainer>
      {/* create a new oauth client */}

      <form class='w-96'>
        <Label for='appName'>App Name</Label>
        <Input name='appName' id='appName' />

        <br />

        <Label for='appLogo'>App Logo</Label>
        <Input name='appLogo' id='appLogo' />
      </form>
    </RootContainer>
  );
};

export default HomePage;
