import React from 'react';
import { config } from './config/gluestack-ui.config';
import { GluestackUIProvider, Text, View } from '@gluestack-ui/themed';

const App = (): React.JSX.Element => {
  return (
    <GluestackUIProvider config={config}>
      <View>
        <Text>Hello world</Text>
      </View>
    </GluestackUIProvider>
  );
};

export default App;
