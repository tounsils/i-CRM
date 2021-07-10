<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Str;
use App\Models\appusers;

class UsersSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        for ($i=0; $i < 6; $i++) { 
	    	appusers::create([
	            'first_name' => Str::random(10),
	            'last_name' => Str::random(10),
	            'email' => Str::random(10).'@gmail.com',
	            'phone' => Str::random(25)
	        ]);
    	}
    }
}

