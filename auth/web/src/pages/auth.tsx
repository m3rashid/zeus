import { Component, Show, createSignal } from 'solid-js';

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
import axios from 'axios';

const Auth: Component = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [errors, setErrors] = createSignal<string[]>([]);

  const currentRoute =
    location.pathname.split('/').at(-1) === 'login' ? 'login' : 'register';

  const onSubmit = async (
    event: Event & {
      currentTarget: HTMLFormElement;
      target: Element;
    }
  ) => {
    event.preventDefault();
    const formErrors: string[] = [];
    const formData = new FormData(event.currentTarget);
    const values = Object.fromEntries(formData);
    if (!values.email || !values.email.toString().includes('@'))
      formErrors.push('Please enter a valid email');
    if (!values.password) formErrors.push('Please enter Password');

    if (currentRoute === 'register') {
      if (!values.confirmPassword)
        formErrors.push('Please confirm your Password');
      if (!values.name) formErrors.push('Please enter your name');
      if (values.password !== values.confirmPassword)
        formErrors.push('Passwords do not match');
    }
    if (formErrors.length > 0) {
      setErrors(formErrors);
      return;
    }

    try {
      await axios.post(
        `http://localhost:5000/api/anonymous/flow/${currentRoute}`,
        values
      );
      navigate('/auth/choose-user' + location.search);
    } catch (err) {
      setErrors(['There was an error contacting the server']);
    }
  };

  return (
    <div class='w-full min-h-screen flex items-center justify-center overflow-auto h-full'>
      <Tabs
        class='w-[400px]'
        defaultValue={currentRoute}
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
                Please enter your credentials to login.
              </CardDescription>
            </CardHeader>
            <form onSubmit={onSubmit}>
              <CardContent class='space-y-2'>
                <div class='space-y-1'>
                  <Label for='email'>Email</Label>
                  <Input
                    name='email'
                    id='email'
                    type='email'
                    placeholder='someone@example.com'
                  />
                </div>
                <div class='space-y-1'>
                  <Label for='password'>Password</Label>
                  <Input
                    name='password'
                    id='password'
                    placeholder='Shh ...'
                    type='password'
                  />
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
              <CardTitle>Register</CardTitle>
              <CardDescription>
                Please enter your details to create account
              </CardDescription>
            </CardHeader>
            <form onSubmit={onSubmit}>
              <CardContent class='space-y-2'>
                <div class='space-y-1'>
                  <Label for='name'>Name</Label>
                  <Input name='name' id='name' placeholder='Someone' />
                </div>
                <div class='space-y-1'>
                  <Label for='email'>Email</Label>
                  <Input
                    id='email'
                    name='email'
                    type='email'
                    placeholder='someone@example.com'
                  />
                </div>
                <div class='space-y-1'>
                  <Label for='password'>Password</Label>
                  <Input
                    name='password'
                    id='password'
                    placeholder='Shh ...'
                    type='password'
                  />
                </div>
                <div class='space-y-1'>
                  <Label for='confirmPassword'>Password</Label>
                  <Input
                    name='confirmPassword'
                    id='confirmPassword'
                    placeholder='Shh ...'
                    type='password'
                  />
                </div>
              </CardContent>
              <CardFooter>
                <Button type='submit'>Create Account</Button>
              </CardFooter>
            </form>
          </Card>
        </TabsContent>

        <Show when={errors().length > 0}>
          <Card class='mt-2'>
            <CardHeader>
              <CardTitle>Errors</CardTitle>
            </CardHeader>
            <CardContent>
              <div class='space-y-1'>
                {errors().map((err) => (
                  <p>{err}</p>
                ))}
              </div>
            </CardContent>
          </Card>
        </Show>
      </Tabs>
    </div>
  );
};

export default Auth;
