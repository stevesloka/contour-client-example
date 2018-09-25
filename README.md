# Contour Client Example

This repo is a very simple example of how to import the Contour client to manage IngressRoute CRD's.

# Sample Output

```sh
$ go run main.go --kubecfg-file=/Users/stevesloka/.kube/aws/kubeconfig                                                                                                          [14:20:28]
Got Routes:  {{IngressRoute contour.heptio.com/v1beta1} {kuard  default /apis/contour.heptio.com/v1beta1/namespaces/default/ingressroutes/kuard 85dade03-c0ee-11e8-a736-06b436d871a6 1023250 1 2018-09-25 14:12:11 -0400 EDT <nil> <nil> map[] map[kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"contour.heptio.com/v1beta1","kind":"IngressRoute","metadata":{"annotations":{},"name":"kuard","namespace":"default"},"spec":{"routes":[{"match":"/","services":[{"name":"node02-kuard","port":80},{"name":"node02-nginx","port":80}]}],"virtualhost":{"fqdn":"marketing.local"}}}
] [] nil [] } {0xc4203f9e60 [{/ [{node02-kuard 80 0 <nil> } {node02-nginx 80 0 <nil> }] { } false false}]} {valid valid IngressRoute}}
```
