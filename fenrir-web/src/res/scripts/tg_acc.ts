import { ref } from 'vue';

export interface Account {
  uid: string;
  id: number;
  first_name: string;
  last_name: string;
  phone: string;
  username: string;
}

export interface Dialog {
  vsname: string;
  id: number;
  username: string;
}

export function useAccounts() {
  const accounts = ref<Account[]>([]);
  const dialogs = ref<Dialog[]>([]);
  const showmsgac = ref(true);
  const showmsgch = ref(true);

  const getAccs = async () => {
    showmsgac.value = false;
    try {
      const response = await fetch('/fenrir/rest', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ action: "get_accounts", data: "" })
      });
      const result = await response.json();
      accounts.value = result.accounts.accounts || result.accounts;
    } catch (error) {
      console.error("Accounts not loaded:", error);
    }
  };

  const getDialogs = async (data: string) => {
    showmsgch.value = false;
    try {
      const response = await fetch('/fenrir/rest', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ action: "get_dialogs", data: data })
      });
      const result = await response.json();
      dialogs.value = result.dialogs || [];
      console.log("Dialogs loaded:", dialogs.value);
    } catch (error) {
      console.error("Dialogs not loaded:", error);
    }
  };

  return { accounts, dialogs, showmsgac, showmsgch, getAccs, getDialogs };
}