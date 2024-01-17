import '@testing-library/jest-dom';
import { describe, expect, it } from 'vitest';
import { render } from '@solidjs/testing-library';

import App from '../src/app';

describe('App', () => {
  it('should render the app', () => {
    const { getByText } = render(() => <App />);
    expect(getByText('Hello')).toBeInTheDocument();
  });
});
