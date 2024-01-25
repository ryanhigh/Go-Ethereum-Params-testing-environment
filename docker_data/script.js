function sendTransaction(){
  eth.sendTransaction({
    from: '1f520a1a350bf6c6eece1af0d23978e34b5956c8',
    to: 'ab9605ed5b7dc5c8283fc4f6d83a34d583e37ec5',
    value: web3.toWei(6, 'ether')
  });
}


// 每隔一段时间发送一次交易
setInterval(sendTransaction, 10); // 5000 毫秒为间隔，即每 5 秒发送一次
