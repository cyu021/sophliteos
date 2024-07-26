import { genMessage } from '../helper';
import antdLocale from 'ant-design-vue/es/locale/ko_KR';

const modules = import.meta.globEager('./ko/**/*.json');
export default {
  message: {
    ...genMessage(modules, 'ko'),
    antdLocale,
  },
  dateLocale: null,
  dateLocaleName: 'ko',
};
