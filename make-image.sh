sudo make build
sudo docker build -t localhost:5000/xpu-sched:v3 .
sudo docker push localhost:5000/xpu-sched:v3
sudo crictl pull localhost:5000/xpu-sched:v3