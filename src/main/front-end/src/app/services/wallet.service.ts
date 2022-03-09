import {Injectable} from '@angular/core';
import * as bip39 from 'bip39';

@Injectable({
  providedIn: 'root'
})
export class WalletService {

  constructor() {
  }

  /**
   * 根据长度生成助记词(len should be 12、15、18、21、24)
   * CS = ENT / 32
   * MS = (ENT + CS) / 11
   *
   * |  ENT  | CS | ENT+CS |  MS  |
   * +-------+----+--------+------+
   * |  128  |  4 |   132  |  12  |
   * |  160  |  5 |   165  |  15  |
   * |  192  |  6 |   198  |  18  |
   * |  224  |  7 |   231  |  21  |
   * |  256  |  8 |   264  |  24  |
   * @param len
   */
  generateMnemonicByLen(len: number): string[] {
    const mnemonic = bip39.generateMnemonic(len * 11 * 32 / 33);
    return mnemonic.split(' ');
  }

  /**
   * 随机助记词
   */
  generateMnemonic(): string[] {
    const mnemonic = bip39.generateMnemonic();
    return mnemonic.split(' ');
  }

  /**
   * 验证助记词的合法性(array)
   * @param mnemonic 助记词数组
   */
  validateMnemonic(mnemonic: string[]) {
    return bip39.validateMnemonic(mnemonic.join(' '));
  }

  testMnemonic() {
    console.log('ge: ', this.generateMnemonic());
    console.log('ge by len: ', this.generateMnemonicByLen(12));
    console.log('ge by len: ', this.generateMnemonicByLen(15));
    console.log('ge by len: ', this.generateMnemonicByLen(12));
    console.log('ge by len: ', this.generateMnemonicByLen(12));
    console.log('validate: ', this.validateMnemonic(['omit', 'sock', 'rail', 'lunch', 'spend', 'rough', 'ship', 'artwork', 'range', 'similar', 'grow', 'velvet']));
  }
}
