import { Button } from '@admin/components/ui/button';
import ModeToggle from '@admin/header/modeToggle';
import { ColorModeProvider } from '@kobalte/core';

const HomePage = () => {
  return (
    <div>
      <ColorModeProvider>
        <div>hello</div>
        <ModeToggle />
        <Button variant='destructive'>Click me</Button>
      </ColorModeProvider>
    </div>
  );
};

export default HomePage;
