<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\appusers;  // to be added

class UsersSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
      // \App\Models\User::factory(10)->create();
      appusers::factory()
      ->count(10)
      ->create();
    }
}

