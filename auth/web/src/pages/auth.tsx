import { Component } from 'solid-js';

import {
  Tabs,
  TabsList,
  TabsTrigger,
  TabsContent,
} from '@web/components/ui/tabs';
import {
  Card,
  CardTitle,
  CardFooter,
  CardHeader,
  CardContent,
  CardDescription,
} from '@web/components/ui/card';
import { Input } from '@web/components/ui/input';
import { Label } from '@web/components/ui/label';
import { Button } from '@web/components/ui/button';
import { useLocation, useNavigate } from '@solidjs/router';

const Auth: Component = () => {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <div class='w-full h-screen flex items-center justify-center'>
      <Tabs
        class='w-[400px]'
        defaultValue={
          location.pathname.split('/').at(-1) === 'login' ? 'login' : 'register'
        }
        onChange={(route) => navigate('/auth/' + route + location.search)}
      >
        <TabsList class='grid w-full grid-cols-2'>
          <TabsTrigger value='login'>Login</TabsTrigger>
          <TabsTrigger value='register'>Register</TabsTrigger>
        </TabsList>

        <TabsContent value='login'>
          <Card>
            <CardHeader>
              <CardTitle>Login</CardTitle>
              <CardDescription>
                Please Enter your credentials to login.
              </CardDescription>
            </CardHeader>
            <form>
              <CardContent class='space-y-2'>
                <div class='space-y-1'>
                  <Label for='email'>Email</Label>
                  <Input id='email' placeholder='someone@example.com' />
                </div>
                <div class='space-y-1'>
                  <Label for='password'>Password</Label>
                  <Input id='password' placeholder='Shh ...' type='password' />
                </div>
              </CardContent>
              <CardFooter>
                <Button type='submit'>Login</Button>
              </CardFooter>
            </form>
          </Card>
        </TabsContent>

        <TabsContent value='register'>
          <Card>
            <CardHeader>
              <CardTitle>Password</CardTitle>
              <CardDescription>
                Change your password here. After saving, you'll be logged out.
              </CardDescription>
            </CardHeader>
            <CardContent class='space-y-2'>
              <div class='space-y-1'>
                <Label for='current'>Current password</Label>
                <Input id='current' type='password' />
              </div>
              <div class='space-y-1'>
                <Label for='new'>New password</Label>
                <Input id='new' type='password' />
              </div>
            </CardContent>
            <CardFooter>
              <Button>Save password</Button>
            </CardFooter>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
};

export default Auth;
