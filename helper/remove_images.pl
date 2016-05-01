#!/usr/bin/perl 

use strict;
use warnings;

sub main {
    my $results = `docker images`;
    my @lines = split /\n/, $results;
    shift @lines;

    my @image_ids = ();

    for my $line ( @lines ) {
        my @comps = split /[\s\t]+/, $line;
        push(@image_ids, $comps[2]);
    }

    for my $image_id ( @image_ids ) {
        print "remove image: $image_id\n";
        my $result = `docker rmi $image_id`;
    }
}

&main();
