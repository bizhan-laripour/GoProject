string="salam delbar,khubi?,mikham begam,kheili duset daram,mikham vasat,sher bekhunam,dust dari,beshnavi?"
sher="hasani nagoo,bala begu,tanbale tanbala,begoo,muye bolan,ruye siah,nakhune deraz,vaho vaho vah,na felfeli,na ghelgheli,na morghe zarde,kakoli,hichkas bahash,refigh nabud,hasani ruye,sepaye,neshaste bud,tu saye,hasani ruye,sepaye,neshaste bud,tu saye,hoooraaa!,hoooraaa!,kheili asheghetam,delbare ghashangam,boos behet."
IFS=',' read -ra array <<< "$string"
for mess in "${array[@]}";
do
figlet $mess
done
echo type kon y/n ro:

read confirmation

if [ $confirmation == y ]; then
 IFS=',' read -ra sherArray <<< "$sher"
  for val in "${sherArray[@]}";
  do
    figlet $val
    done
    fi