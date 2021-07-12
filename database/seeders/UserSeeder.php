<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\user;  // to be added

class UserSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
      // \App\Models\User::factory(10)->create();
      user::factory()
      ->count(3)
      ->create();
    }
}

