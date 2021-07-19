import React from 'react';
import { hot } from 'react-hot-loader/root';
import { Button, Card } from 'antd';
import './app.less';

const App = () => {
    return <div style={{padding: 20, color: 'var(--primary-color)'}}>
        Hello
        <Button type='primary'>123</Button>
        <Card>
            This is a dynamic theme.
        </Card>
    </div>
}

export default hot(App);