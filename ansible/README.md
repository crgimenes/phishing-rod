

```bash
ansible-playbook -i inventory.ini upgrade_all.yaml
```

Use ansible-lint:

```bash
ansible-lint --write
```




TODO: add to ansible 

```bash
wget -O - https://www.kismetwireless.net/repos/kismet-release.gpg.key | sudo apt-key add -
echo "deb https://www.kismetwireless.net/repos/apt/release/$(lsb_release -cs) $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/kismet.list
apt update
apt install kismet
```


