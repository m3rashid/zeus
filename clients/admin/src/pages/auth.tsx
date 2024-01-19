import { Button } from '@admin/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@admin/components/ui/card';
import { Input } from '@admin/components/ui/input';
import { Label } from '@admin/components/ui/label';
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from '@admin/components/ui/tabs';
import { Component } from 'solid-js';

const Auth: Component = () => {
  return (
    <div class='w-full h-screen flex items-center justify-center'>
      <Tabs defaultValue='account' class='w-[400px]'>
        <TabsList class='grid w-full grid-cols-2'>
          <TabsTrigger value='login'>Login</TabsTrigger>
          <TabsTrigger value='register'>Register</TabsTrigger>
        </TabsList>

        <TabsContent value='login'>
          <Card>
            <CardHeader>
              <CardTitle>Login</CardTitle>
              <CardDescription>
                Make changes to your account here. Click save when you're done.
              </CardDescription>
            </CardHeader>
            <CardContent class='space-y-2'>
              <div class='space-y-1'>
                <Label for='name'>Name</Label>
                <Input id='name' value='Pedro Duarte' />
              </div>
              <div class='space-y-1'>
                <Label for='username'>Username</Label>
                <Input id='username' value='@peduarte' />
              </div>
            </CardContent>
            <CardFooter>
              <Button>Save changes</Button>
            </CardFooter>
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
