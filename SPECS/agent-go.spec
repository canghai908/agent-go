Summary: A zabbix agent written by golang
Name: agent-go
Version: 0.0.1
Release: 1%{?dist}
Source0: %{name}-%{version}.tar.gz
License: GPL
Group: Applications/Internet

%if 0%{?fedora} || 0%{?rhel} == 7
BuildRequires: systemd
Requires(post): systemd
Requires(preun): systemd
Requires(postun): systemd
%endif

%description
Zabbix　Agent　write by Golang
%prep

%build

%install

mkdir -p ${RPM_BUILD_ROOT}/usr/bin/
cp -f /opt/agent-go/agent-go  ${RPM_BUILD_ROOT}/usr/bin/agent-go

mkdir -p ${RPM_BUILD_ROOT}/etc/agent-go/
cp -f /opt/agent-go/conf/agent-go.ini ${RPM_BUILD_ROOT}/etc/agent-go/agent-go.ini

mkdir -p ${RPM_BUILD_ROOT}/lib/systemd/system/
cp -f /opt/agent-go/systemd/agent-go.service ${RPM_BUILD_ROOT}/lib/systemd/system/agent-go.service

%files
%defattr(-,root,root)
/usr/bin/agent-go
/lib/systemd/system/agent-go.service
%dir
/etc/agent-go