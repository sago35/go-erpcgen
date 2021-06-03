use strict;
use warnings;
use utf8;
use Term::Encoding qw(term_encoding);
my $encoding = term_encoding;

binmode STDOUT => "encoding($encoding)";
binmode STDERR => "encoding($encoding)";

use Getopt::Kingpin;

my $kingpin = Getopt::Kingpin->new;
my $decode = $kingpin->flag("decode", "")->bool;
$kingpin->parse;

my @lines = <STDIN>;
chomp @lines;

my $m = get_service_map("all.txt");

# Adding one more line, but not processing the last line added
push @lines, "";
for (my $i = 0; $i < @lines - 1; $i++) {
    if ($lines[$i] =~ /^[-+=]/) {
        # skip
        next;
    } elsif ($lines[$i] =~ /^tx :  4 :/ and $lines[$i+1] =~ /^tx : *\d+ : (.*)/) {
        my ($msgType, $requestNumber, $service, $xVersion) = map {hex} split / /, $1;
        #printf "get : %s\n", $1;
        #printf "    : %02X %02X %02X %02X\n", $msgType, $requestNumber, $service, $xVersion;
        #printf "    : %s %s\n", $m->{$service}->{name}, $m->{$service}->{methods}->{$requestNumber}->{name};
        my $name = $m->{$service}->{methods}->{$requestNumber}->{name} // "not-defined";
        $name =~ s/^.*_rpc/rpc/;
        printf "= %s()\n", $name;
        #use Data::Dumper;
        #print Dumper $m;
    }
    printf "%s\n", $lines[$i];

    if ($decode && $lines[$i] =~ /^[tr]x : *\d+ : (.*)/) {
        my @bytes = map {
            my $c = hex $_;
            if (0x20 <= $c and $c < 0x7F) {
                sprintf "%c", hex $_;
            } else {
                ".";
            }
        } split / /, $1;
        printf "   : [%s]\n", join "", @bytes;
    }
}

use Path::Tiny;
sub get_service_map {
    my $file = shift;
    my @lines = path($file)->lines({chomp => 1});

    my $ret = {};
    my $sname = "";
    my $sid = 0;
    for (my $i = 0; $i < @lines; $i++) {
        if ($lines[$i] =~ /\bkrpc_(\S+)_id = (\d+),$/ and $lines[$i-1] eq "{") {
            $sname = $1;
            $sid = int $2;
            $ret->{$sid} = {
                name    => $sname,
                id      => $sid,
                methods => {},
            };
        } elsif ($lines[$i] =~ /\bkrpc_(\S+)_id = (\d+),$/) {
            $ret->{$sid}->{methods}->{int $2} = {
                name => $1,
                id   => int $2,
            };
        }
    }
    return $ret;
}
