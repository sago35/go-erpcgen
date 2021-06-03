use strict;
use warnings;
use utf8;

use Path::Tiny;

my $p = path("testdata/erpc_idl/rpc_system.erpc");

my $idl = parseIDL($p->parent->absolute, $p->basename);

sub parseIDL {
    my $dir = shift;
    my $basename = shift;
    printf "// %s\n", $dir->child($basename)->relative;

    for my $line ($dir->child($basename)->lines({chomp => 1})) {
        if ($line =~ /^import +"([^"]+)"/) {
            printf "//%s\n", $line;
            my $target = $dir->child($1);
            parseIDL($target->parent->absolute, $target->basename);
        } else {
            printf "%s\n", $line;
        }
    }
}
