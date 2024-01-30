import { defineStore } from 'pinia';
import { AbilityItem } from '/@/api/task/model';
import { getAbilites } from '/@/api/task';

interface AbilityState {
  abilities: AbilityItem[];
  abilitiesMap: Map<number, string>;
}

export const useAbilitesStore = defineStore({
  id: 'abilities',
  state: (): AbilityState => ({
    abilities: [],
    abilitiesMap: new Map(),
  }),
  getters: {
    getAbilities(state): AbilityItem[] {
      return state.abilities;
    },
  },
  actions: {
    async setAbilities() {
      const res = await getAbilites();
      res.forEach((item) => {
        this.abilitiesMap.set(item.type, item.name);
      });
      this.abilities = res.sort((a, b) => a.type - b.type) || [];
      return this.abilitiesMap;
    },
  },
});
