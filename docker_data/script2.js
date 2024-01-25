function sendTransaction(){
  eth.sendTransaction({
    from: 'ab9605ed5b7dc5c8283fc4f6d83a34d583e37ec5',
    to: 'bfa8a5a4476ef8b92decdbf952dc8c3f2a1c1943',
    value: web3.toWei(6, 'ether')
  });
}


// 每隔一段时间发送一次交易
setInterval(sendTransaction, 10); // 5000 毫秒为间隔，即每 5 秒发送一次
