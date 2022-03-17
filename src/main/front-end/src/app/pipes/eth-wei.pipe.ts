import { Pipe, PipeTransform } from '@angular/core';
import Web3 from 'web3'

@Pipe({
  name: 'ethWei'
})
export class EthWeiPipe implements PipeTransform {

  transform(value: string, to: 'eth' | 'wei'): string {
    if (to == 'wei') {
      return Web3.utils.toWei(value, 'wei')
    } else if (to == 'eth') {
      return Web3.utils.fromWei(value, 'ether')
    }
    return "error"
  }
}
