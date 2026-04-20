<style>
    @import "@/res/css/Telegram.css";
</style>

<script setup lang="ts">
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useAccounts, type Account, type Dialog } from '@/res/scripts/tg_acc';

  const OpenChats = ref(false);
  const OpenAccounts = ref(false);
  const activePage = ref('info');
  const router = useRouter();
  const { accounts, dialogs, showmsgac, showmsgch, getAccs, getDialogs } = useAccounts();

  const cur_account = ref<Account>({} as Account);
  const cur_chat = ref<Dialog>({} as Dialog);

  const gDgs = async (account: Account) => {
    await getDialogs(JSON.stringify({ ...account, id: String(account.id) }));
  };

  const toSettings = () => router.push('/settings');
</script>

<template>
  <div class="telegram-route">
    <div class="fenrir-page">
      <dev v-if="activePage === 'info'">
        <span class="cause-text">Information Page</span>
        
        <p class="quicksand-small-text">
          Current Account: {{ cur_account?.first_name + " " + cur_account?.last_name + " (@" + cur_account?.username + ")"}}
        </p>

        <p class="quicksand-small-text">
          Current Chat: {{ cur_chat?.vsname + " (@" + cur_chat?.username + ")" }}
        </p>

        <p>Docs</p>
      </dev>
      <dev v-if="activePage === 'view'">
        <span class="cause-text">View Page>> {{ cur_chat?.vsname + "(" + cur_chat?.id + ")" + "#" + cur_account?.first_name }}</span>


      </dev>
    </div>

    <aside class="fenrir">
      <button class="fenrir-button" @click="activePage = 'info'">Info</button>
      <button class="fenrir-button" @click="activePage = 'view'">View#CHT</button>
      <button class="fenrir-button" @click="OpenChats = true">Chats</button> 
      <button class="fenrir-button" @click="OpenAccounts = true">Accounts</button>

      <button class="set-button" @click="toSettings">⚙️</button>
    </aside>

    <aside :class="['aside', { 'aside-open': OpenChats }]">
        <dev class="aside-head">
          <button @click="OpenChats = false">🢀</button>
          <button @click="gDgs(cur_account)">G</button>
          <span class="cause-text">Fenrir Chats</span>
        </dev>
        <span class="quicksand-small-text" v-if="showmsgch === true">
          To load chats, you need to press the "G" button. To join a chat, you need to use an account in the "Accounts" section.
          <small>Read Info page for more information.</small>
        </span>

        <ul>
         <li class="lichat" v-for="diag in dialogs" :key="diag.id">
            <small class="amaticb-very-small-text">{{ diag.vsname }}</small>
            <small class="amaticb-very-small-text">{{ "@" +diag.username }}</small>

            <small class="amaticb-very-small-text">{{ diag.id }}</small> 

            <button @click="cur_chat = diag">Target</button> 
         </li>
        </ul>
    </aside>

    <aside :class="['aside', { 'aside-open': OpenAccounts }]">
        <dev class="aside-head">
          <button @click="OpenAccounts = false">🢀</button>
          <button @click="getAccs">U</button>
          <span class="cause-text">Fenrir Accs </span>
        </dev>
        <span class="quicksand-small-text" v-if="showmsgac === true">To load the .sessions file, you need to press the "U" button! <small>Read Info page for more information.</small></span>

        <ul>
          <li class="liacc" v-for="acc in accounts" :key="acc.id">
            <small class="amaticb-very-small-text">{{ acc.first_name }} {{ acc.last_name }}</small>
            <small class="amaticb-very-small-text">{{ "@" + acc.username }} {{ "+" + acc.phone.slice(0, -4) + "xxxx"}}</small>
            <button @click="cur_account = acc">Use</button>
          </li>
        </ul>
    </aside>
    
  </div>
</template>
