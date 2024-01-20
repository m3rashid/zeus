import { Button } from '@admin/components/ui/button';
import { Component } from 'solid-js';
import ProjectSelector from './projectSelector';

const SecondaryHeader: Component = () => {
  return (
    <div class='w-full h-9 px-2 border-b-2 bg-slate-200 dark:bg-slate-900 flex items-center'>
      {/* project selector */}
      <ProjectSelector
        trigger={
          <Button class='h-8 w-96' variant='secondary'>
            Project Name
          </Button>
        }
      />
    </div>
  );
};

export default SecondaryHeader;
