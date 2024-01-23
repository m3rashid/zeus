import { Component, createResource } from 'solid-js';

import {
  Card,
  CardTitle,
  CardHeader,
  CardContent,
  CardDescription,
} from '@web/components/ui/card';
import { User } from '@web/lib/types';

const ChooseUser: Component = () => {
  const [users] = createResource<User[]>(
    async () => {
      const res = await fetch('/api/flow/get-select-users');
      if (!res.ok) return [];

      const users = await res.json();
      return users;
    },
    { initialValue: [] }
  );

  return (
    <div class='w-full min-h-screen flex items-center justify-center overflow-auto h-full'>
      <Card class='w-[400px]'>
        <CardHeader>
          <CardTitle>Choose User</CardTitle>
          <CardDescription>
            Choose any one from the given profiles
          </CardDescription>
        </CardHeader>

        <CardContent>
          {users().map((user) => (
            <div class='flex items-center space-x-2'>
              <div class='w-10 h-10 rounded-full bg-gray-200'>
                <img src={user.profilePicUrl} alt={user.name} class='' />
              </div>
              <div class='flex flex-col'>
                <span class='font-semibold'>{user.name}</span>
                <span class='text-sm text-gray-500'>{user.email}</span>
              </div>
            </div>
          ))}
        </CardContent>
      </Card>
    </div>
  );
};

export default ChooseUser;
