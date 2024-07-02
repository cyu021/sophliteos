import { genMessage } from '../helper';
import antdLocale from 'ant-design-vue/es/locale/ja_JP';

const modules = import.meta.globEager('./ja/**/*.json');
export default {
  message: {
    ...genMessage(modules, 'ja'),
    antdLocale,
  },
};
