
令牌：每次都需要拿到令牌后，才可以访问
桶：有一定大小的桶，桶中最多可以放一定数量的令牌
放入频率：按照一定的频率向通里面放入令牌，但是令牌数量不能超过桶的容量


Limter提供三种消费 token 的方式, 用户可以每次消费一个Token，也可以一次性消费多个Token, 而每种方法代表了当Token不足时，各自不同的对应手段。
- Allow / AllowN
  Allow 实际上就是 AllowN(time.Now(),1)。
  AllowN 方法表示，截止到某一时刻，目前桶中数目是否至少为n个，满足则返回true，同时从桶中消费n个token。反之返回不消费Token，返回 false。
  通常对应这样的线上场景: 如果请求速率过快，就直接丢到某些请求。
- Reserve / ReserveN
  Reserve 相当于 ReserveN(time.Now(), 1)。
  ReserveN 的用法就相对来说复杂一些，当调用完成后，无论Token是否充足，都会返回一个 Reservation 对象。
  你可以调用该对象的 Delay()方法，该方法返回了需要等待的时间。如果等待时间为0，则说明不用等待。必须等到等待时间之后，才能进行接下来的工作。或者，如果不想等待，可以调用 Cancel() 方法，该方法会将 Token 归还。
  ```go
  r := lim.Reserve()
  f !r.OK() {
      // Not allowed to act! Did you remember to set lim.burst to be > 0 ?
      return
  }
  time.Sleep(r.Delay())
  Act() // 执行相关逻辑
  ```
- Wait / WaitN
  Wait 实际上就是 WaitN(ctx,1)。
  当使用Wait方法消费Token时，如果此时桶内Token数组不足(小于N)，那么Wait方法将会阻塞一段时间，直至Token满足条件。如果充足则直接返回。
  我们可以设置context的Deadline或者Timeout，来决定此次Wait的最长时间。
  大部分时候使用Wait。

- Limiter 支持可以调整速率和桶大小：
  SetLimit(Limit) 改变放入Token的速率
  SetBurst(int) 改变Token桶大小