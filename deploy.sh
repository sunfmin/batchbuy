goxc -bc 'linux'
scp -p222 batchbuy_linux_amd64.tar.gz lowtea@115.238.41.246:~/batchbuy_linux_amd64.tar.gz
scp -P222 run_on_cn.sh lowtea@115.238.41.246:/home/lowtea
ssh -p222 lowtea@115.238.41.246 "chmod +x /home/lowtea/run_on_cn.sh && sudo /home/lowtea/run_on_cn.sh"
