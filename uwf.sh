for svc in ssh http https
do 
ufw allow $svc
done
