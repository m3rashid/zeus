import { Component, createResource } from 'solid-js';

import {
  Card,
  CardTitle,
  CardHeader,
  CardContent,
  CardDescription,
} from '@web/components/ui/card';
import { User } from '@web/lib/types';
import axios from 'axios';

const ChooseUser: Component = () => {
  const [users] = createResource<User[]>(
    async () => {
      try {
        const { data } = await axios.get('/api/flow/users');
        console.log(data);
        return data;
      } catch (err) {
        return [];
      }
    },
    { initialValue: [] }
  );

  const handleSelectUser = async (userId: string) => {
    try {
      const { data } = await axios.post(
        'http://localhost:5000/api/flow/select-user',
        { userId: Number(userId) }
      );
      console.log(data);
    } catch (err) {
      console.log(err);
    }
  };

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
          <div class='space-y-4'>
            {users().map((user) => (
              <div
                class='flex items-center space-x-2 cursor-pointer'
                onClick={() => handleSelectUser(user.id)}
              >
                <div class='w-10 h-10 rounded-full bg-gray-200'>
                  {user.profilePicUrl ? (
                    <img src={user.profilePicUrl} alt={user.name} class='' />
                  ) : null}
                </div>
                <div class='flex flex-col'>
                  <span class='font-semibold'>{user.name}</span>
                  <span class='text-sm text-gray-500'>{user.email}</span>
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default ChooseUser;
