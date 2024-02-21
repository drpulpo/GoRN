import { StatusBar } from 'expo-status-bar';
import { Button, StyleSheet, Text, View } from 'react-native';
import { Header } from './components/Header';
import { Footer } from './components/Footer';
import { OrderList } from './components/home/OrderList';
import { useOrders } from './hooks/useOrders';

export default function App() {

  const { orders } = useOrders()
  return (
    <View style={styles.container}>
      <StatusBar style="auto" />
      <Header />
      <main>
        <OrderList orders={orders}></OrderList>
      </main>
      <Footer />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
