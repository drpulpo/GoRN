import { StatusBar } from 'expo-status-bar';
import { Button, StyleSheet, View } from 'react-native';
import { Header } from './components/Header';
import { Footer } from './components/Footer';
import { OrderList } from './components/home/OrderList';
import { useOrders } from './hooks/useOrders';
import { OrderForm } from './components/home/OrderForm';

export default function App() {
  const { orders, selectedOrder, createOrder, saveOrder, cancelSaveOrder } = useOrders()

  const handleNewOrderPress = () => {
    createOrder()
  }

  const handleSaveOrderForm = () => {
    saveOrder()
  }

  const handleCancelOrderForm = () => {
    cancelSaveOrder()
  }

  return (
    <View style={styles.container}>
      <StatusBar style="auto" />
      <Header />
      <View style={styles.main}>
        <View style={styles.mainContainer}>
          <View style={styles.spacer}>
          </View>
          {selectedOrder.Id ?
            <View style={styles.newOrderContainer}>
              <OrderForm selectedOrder={selectedOrder} onSave={handleSaveOrderForm} onCancel={handleCancelOrderForm} />
            </View> :
            <View style={styles.newOrderContainer}>
              <Button color={styles.newOrderButton.backgroundColor} title='+ New Order' onPress={handleNewOrderPress}></Button>
              <OrderList orders={orders}></OrderList>
            </View>
          }
          <View style={styles.spacer}>
          </View>
        </View>
      </View>
      <Footer />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    display: 'flex',
    flex: 1,
    backgroundColor: '#CFC',
    alignItems: 'center',
    justifyContent: 'center',
  },
  main: {
    width: '100%',
    flex: 1,
    margin: 30
  },
  mainContainer: {
    flexDirection: 'row'
  },
  newOrderButton: {
    backgroundColor: '#21d321'
  },
  spacer: {
    flex: 1
  },
  newOrderContainer: {
    flex: 10,
    flexDirection: 'column'
  }
});
