import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTrigger,
} from '@admin/components/ui/dialog';
import { Component, JSX } from 'solid-js';

export type ProjectSelectorProps = {
  trigger: JSX.Element;
};

const ProjectSelector: Component<ProjectSelectorProps> = (props) => {
  return (
    <Dialog modal>
      <DialogTrigger>
        {props.trigger}
        <span class='sr-only'>Project Details</span>
      </DialogTrigger>

      <DialogContent class='top-20'>
        {/* last tab to create new project */}
        <DialogHeader>Select Project</DialogHeader>

        {/* tab view for multiple projects */}
      </DialogContent>
    </Dialog>
  );
};

export default ProjectSelector;
