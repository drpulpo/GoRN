import { Image, Text, View, StyleSheet } from 'react-native'

export const Header = () => {
    return (
        <View style={styles.container}>
            <View style={styles.leftHeader}>
                {/* <Image></Image> */}
                <Text style={styles.brandName}>GoRN</Text>
            </View>
            {/* <View style={styles.rightHeader}>
                <Text style={styles.menu}>Menu</Text>
            </View> */}
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        width: '100%',
        display: 'flex',
        flexDirection: 'row',
        height: 50,
        backgroundColor: '#21d321',
        justifyContent: 'space-between',
        paddingTop: 10,
        paddingLeft: 50,
        color: '#FFFFFF',
        fontSize: 20
    },
    leftHeader: {
        display: 'flex',
        flexDirection: 'row'
    },
    brandName: {
        color: '#FFFFFF',
        fontSize: 20
    },
    rightHeader: {
        display: 'flex'
    },
    menu: {
        color: '#FFFFFF',
        fontSize: 20
    }
})