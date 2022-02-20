
From [this](https://docs.docker.com/compose/gettingstarted/) tutorial, slightly
modified (exposed the app on port 5050 instead of 5000).

```
$ docker-compose build

$ docker image ls
REPOSITORY                 TAG                 IMAGE ID            CREATED             SIZE
compose-tutorial_web       latest              d0fde14f6baa        56 seconds ago      183MB
python                     3.7-alpine          4d1c95b3db1c        8 days ago          41.9MB
yandex/clickhouse-server   20                  633f979d7ed9        11 months ago       809MB
zookeeper                  latest              82f175cb301d        11 months ago       253MB
zookeeper                  3.5                 bc94c000b74c        11 months ago       245MB
busybox                    latest              f0b02e9d092d        12 months ago       1.23MB
harveyos/harvey            latest              2ab13768446b        15 months ago       1.53GB
hello-world                latest              bf756fb1ae65        22 months ago       13.3kB
elasticsearch              6.5.1               32f93c89076d        2 years ago         773MB

$ docker save -o ./tutorial-image.tar compose-tutorial_web

(then scp tutorial-image.tar to target, then `docker load -i tutorial-image.tar`
there)
```

Workable ffmpeg video rotation:

$ time ffmpeg -i ~/Downloads/portrait-video.mp4 -vf "transpose=2,pad=width=1920:height=1080:x=0:y=540:color=black" output2.mp4

composed from these stack answers:
https://superuser.com/questions/690021/video-padding-using-ffmpeg/690211
https://stackoverflow.com/questions/3937387/rotating-videos-with-ffmpeg/18624384

This one doesn't do an unnecessary transposing, just makes the video big enough
to include the portrait. Still need to change x or y param to move it into the
middle of the landscape frame.

$ sudo time ffmpeg -i /volume1/video/VID_20211017_131432.mp4 -vf "pad=width=3412:height=1920:x=0:y=120:color=black" output3.mp4


This one works faster (due to '-preset veryfast' and '-codec:a copy', which
copies the audio stream verbatim), and should move the image to the center.

$ sudo time ffmpeg -i /volume1/video/VID_20211017_131432.mp4 -vf "pad=width=3412:height=1920:x=0:y=1166:color=black" -preset veryfast -codec:a copy output4.mp4

`-nostats` flag makes the output friendlier for redirection


Properly centered:

$ sudo time ffmpeg -i /volume1/video/VID_20211017_131432.mp4 -vf "pad=width=3412:height=1920:x=1166:y=0:color=black" -preset veryfast -codec:a copy output5.mp4

