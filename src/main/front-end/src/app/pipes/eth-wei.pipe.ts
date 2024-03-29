import { Pipe, PipeTransform } from '@angular/core';
import Web3 from 'web3'

@Pipe({
  name: 'ethWei'
})
export class EthWeiPipe implements PipeTransform {

  transform(value: string, to: 'eth' | 'wei' | 'gwei' | 'wei2gwei'): string {
    try {
      if (to == 'wei') {
        return Web3.utils.toWei(value, 'ether') // eth => wei
      } else if (to == 'eth') {
        return Web3.utils.fromWei(value, 'ether') // wei => eth
      } else if (to == 'gwei') {
        return Web3.utils.fromWei(Web3.utils.toWei(value, 'gwei')) // gwei => eth
      } else if (to == 'wei2gwei') {
        return Web3.utils.fromWei(value, 'gwei') // wei => gwei
      }
      return "error"
    } catch (e) {
      return "error"
    }
  }
}
