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
      const res = await fetch('/api/users');
      if (!res.ok) return [];

      const users = await res.json();
      return users;
    },
    { initialValue: [] }
  );

  return (
    <Card>
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
              <img src={user.profilePictureUrl} alt={user.name} class='' />
            </div>
            <div class='flex flex-col'>
              <span class='font-semibold'>{user.name}</span>
              <span class='text-sm text-gray-500'>{user.email}</span>
            </div>
          </div>
        ))}
      </CardContent>
    </Card>
  );
};

export default ChooseUser;
